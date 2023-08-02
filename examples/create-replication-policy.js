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
    const now = Date.now()

    const params = {
        name: `harbor-${now}`,
        url: getEnv('HARBOR_REGISTRY_URL', 'https://demo.goharbor.io'),
        type: 'harbor',
        credential: {
            accessKey: getEnv('HARBOR_REGISTRY_USERNAME', 'admin'),
            accessSecret: getEnv('HARBOR_REGISTRY_PASSWORD', 'Harbor12345'),
            type: "basic"
        },
    }

    const registryId = harbor.createRegistry(params)

    return {
        now,
        registryId,
    }
}

export default function ({ now, registryId }) {
    const i = counter.up() - 1

    const params = {
        deletion: false,
        destRegistry: { id: registryId },
        enabled: true,
        filters: [
            { type: "name", value: "library/**" },
            { type: "tag", value: "**" },
        ],
        name: `policy-${now}-${i}`,
        override: true,
        trigger: { type: "manual" },
    }

    try {
        harbor.createReplicationPolicy(params)
    } catch (e) {
        console.log(e)
        errorRate.add(true)
    }
}

export function teardown({ now, registryId }) {
    if (teardownResources) {

        const pageSize = 50

        while (true) {
            try {
                const { policies } = harbor.listReplicationPolicies({ name: `policy-${now}-`, pageSize })

                for (const policy of policies) {
                    try {
                        harbor.deleteReplicationPolicy(policy.id)
                    } catch (e) {
                        console.log(e)
                    }
                }

                if (policies.length < pageSize) {
                    break
                }
            } catch (e) {
                console.log(e)
                break
            }
        }

        try {
            harbor.deleteRegistry(registryId)
        } catch (e) {
            console.log(e)
        }
    }
}
