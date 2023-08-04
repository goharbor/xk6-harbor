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

const teardownResources = getEnv('TEARDOWN_RESOURCES', 'true') === 'true'

export let successRate = new Rate('success');

export let options = {
    setupTimeout: '1h',
    teardownTimeout: '30m',
    noUsageReport: true,
    vus: 10,
    iterations: 20,
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
    const projectName = `project-${Date.now()}`

    harbor.createProject({ projectName })

    console.log(`project ${projectName} created`)

    return {
        projectName,
    }
}

export default function ({ projectName }) {
    const i = counter.up() - 1

    try {
        harbor.createRobot({
            name: `robot-${i}`,
            level: "project",
            permissions: [
                { namespace: projectName, "kind": "project", access: [{ resource: "repository", action: "pull" }] },
            ]
        })
        successRate.add(true)
    } catch (e) {
        console.log(e)
        successRate.add(false)
    }
}

export function teardown({ projectName }) {
    if (teardownResources) {
        const pageSize = 15
    }
}
