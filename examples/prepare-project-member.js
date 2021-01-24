// Create 20 project members for each project which prepared before

import counter from 'k6/x/counter'
import harbor from 'k6/x/harbor'
import { Rate } from 'k6/metrics'

const missing = Object()

function getEnv(env, def = missing) {
    const value = __ENV[env] ? __ENV[env] : def
    if (value === missing) {
        throw (`${env} envirument is required`)
    }

    return value
}

function getEnvInt(env, def = missing) {
    return parseInt(getEnv(env, def), 10)
}

const usersPerProject = getEnvInt('PREPARE_USERS_PER_PROJECT', '20')
const projectsCount = getEnvInt('PREPARE_PROJECTS_COUNT', '500')

export let errorRate = new Rate('errors');

export let options = {
    noUsageReport: true,
    vus: usersPerProject,
    iterations: projectsCount * usersPerProject,
    thresholds: {
        'iteration_duration{scenario:default}': [
            `max>=0`,
        ],
        'iteration_duration{group:::setup}': [`max>=0`],
        'iteration_duration{group:::teardown}': [`max>=0`]
    }
};

export function setup() {
    harbor.initialize({
        scheme: getEnv('HARBOR_SCHEME', 'https'),
        host: getEnv('HARBOR_HOST'),
        username: getEnv('HARBOR_USERNAME', 'admin'),
        password: getEnv('HARBOR_PASSWORD', 'Harbor12345'),
        insecure: true,
    })

    let page = 1
    const pageSize = 100

    const projectIDs = []
    while (projectIDs.length < projectsCount) {
        const { projects } = harbor.listProjects({ page, pageSize })

        for (const project of projects) {
            projectIDs.push(project.projectID)
        }

        if (projects.length == 0 || projects.length < pageSize) {
            break
        }

        page++
    }

    return {
        projectIDs: projectIDs.slice(0, projectsCount)
    }
}

export default function ({ projectIDs }) {
    const i = counter.up()

    let userID = null
    try {
        userID = harbor.createUser(`user-${Date.now()}-${i}`)
    } catch (e) {
        console.log(e)
        errorRate.add(true)

        return
    }

    const projectID = projectIDs[(i - 1) % projectIDs.length]
    try {
        harbor.createProjectMember(projectID, userID)
    } catch (e) {
        console.log(e)
        errorRate.add(true)
    }
}

export function teardown() {
}
