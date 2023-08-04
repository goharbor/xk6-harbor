// Create 20 project members for each project which prepared before

import counter from 'k6/x/counter'
import { Harbor } from 'k6/x/harbor'
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

const harbor = new Harbor({
    scheme: getEnv('HARBOR_SCHEME', 'https'),
    host: getEnv('HARBOR_HOST'),
    username: getEnv('HARBOR_USERNAME', 'admin'),
    password: getEnv('HARBOR_PASSWORD', 'Harbor12345'),
    insecure: true,
})

export function setup() {
    let page = 1
    const pageSize = 100

    const projectNames = []
    while (projectNames.length < projectsCount) {
        const { projects } = harbor.listProjects({ page, pageSize })

        for (const project of projects) {
            projectNames.push(project.name)
        }

        if (projects.length == 0 || projects.length < pageSize) {
            break
        }

        page++
    }

    return {
        projectNames: projectNames.slice(0, projectsCount)
    }
}

export default function ({ projectNames }) {
    const i = counter.up()

    let userID = null
    try {
        userID = harbor.createUser(`user-${Date.now()}-${i}`)
    } catch (e) {
        console.log(e)
        errorRate.add(true)

        return
    }

    const projectName = projectNames[(i - 1) % projectNames.length]
    try {
        harbor.createProjectMember(projectName, userID)
    } catch (e) {
        console.log(e)
        errorRate.add(true)
    }
}

export function teardown() {
}
