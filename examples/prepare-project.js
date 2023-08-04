// Remove the library project and create 500 projects

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

const projectsCount = getEnvInt('PREPARE_PROJECTS_COUNT', '500')

export let errorRate = new Rate('errors');

export let options = {
    noUsageReport: true,
    vus: 100,
    iterations: projectsCount,
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
    try {
        harbor.deleteProject('library')
    } catch (e) {
        console.log(e)
    }
}

export default function () {
    const projectName = `project-${Date.now()}-${counter.up()}`

    try {
        harbor.createProject({ projectName })
    } catch (e) {
        console.log(e)
        errorRate.add(true)
    }
}

export function teardown() {
}
