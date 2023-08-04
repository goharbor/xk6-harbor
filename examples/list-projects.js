// List projects, will prepare 10k projects

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
    setupTimeout: '24h',
    teardownTimeout: '1h',
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
    const { total } = harbor.listProjects({ page: 1, pageSize: 1 })

    console.log(`total projects: ${total}`)

    return {
        projectsCount: total
    }
}

export default function ({ projectsCount }) {
    const pageSize = 15
    const pages = Math.ceil(projectsCount / pageSize)
    const page = Math.floor(Math.random() * pages) + 1

    try {
        harbor.listProjects({ page, pageSize })
    } catch (e) {
        console.log(e)
        errorRate.add(true)
    }
}

export function teardown() {
}
