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
    return {
        now: Date.now()
    }
}

export default function ({ now }) {
    const i = counter.up() - 1

    const params = {
        name: `harbor-${now}-${i}`,
        url: getEnv('HARBOR_REGISTRY_URL', 'https://demo.goharbor.io'),
        type: 'harbor',
        credential: {
            accessKey: getEnv('HARBOR_REGISTRY_USERNAME', 'admin'),
            accessSecret: getEnv('HARBOR_REGISTRY_PASSWORD', 'Harbor12345'),
            type: "basic"
        },
    }

    try {
        harbor.createRegistry(params)
    } catch (e) {
        console.log(e)
        errorRate.add(true)
    }
}

export function teardown({ now }) {
    if (teardownResources) {
        while (true) {
            try {
                const { registries } = harbor.listRegistries({ q: `name=~harbor-${now}-` })
                for (const registry of registries) {
                    harbor.deleteRegistry(registry.id)
                }
            } catch (e) {
                console.log(e)
            }
        }

    }
}
