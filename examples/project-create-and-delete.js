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

export let errorRate = new Rate('errors');

export let options = {
    setupTimeout: '1h',
    teardownTimeout: '30m',
    noUsageReport: true,
    vus: 500,
    iterations: 1000,
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
}

export default function () {
    const projectName = `project-${Date.now()}-${__VU}-${__ITER}`

    try {
        harbor.createProject({ projectName })
        harbor.deleteProject(projectName)
    } catch (e) {
        console.log(e)
        errorRate.add(true)
    }
}

export function teardown() {
}
