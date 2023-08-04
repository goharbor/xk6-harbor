import counter from 'k6/x/counter'
import { Harbor, ContentStore } from 'k6/x/harbor'
import { SharedArray } from 'k6/data'
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

const teardownResources = getEnv('TEARDOWN_RESOURCES', 'true') === 'true'

const harborRegistryURLs = getEnv('HARBOR_REGISTRY_URLS').split(',')
const artifactsCount = getEnvInt('ARTIFACTS_COUNT', '1000')
const artifactSize = getEnv('ARTIFACT_SIZE', '1 KiB')
const policiesCount = getEnvInt('POLICIES_COUNT', '1000')

const contentStore = new ContentStore('replication')

let contents = new SharedArray('contents', function () {
    return contentStore.generateMany(artifactSize, artifactsCount)
});

export let errorRate = new Rate('errors');

export let options = {
    setupTimeout: '1h',
    duration: '2h',
    teardownTimeout: '30m',
    noUsageReport: true,
    vus: 500,
    iterations: 500,
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

    const projectName = `project-${now}`
    harbor.createProject({ projectName })
    console.log(`project ${projectName} created`)

    for (let i = 0; i < artifactsCount; i++) {
        const blob = contents[i % contents.length]
        harbor.push({
            ref: `${projectName}/repository-${now}:tag-${i}`,
            store: contentStore,
            blobs: [blob],
        })
    }

    const registryIds = []
    for (let i = 0; i < harborRegistryURLs.length; i++) {
        const params = {
            name: `harbor-${now}-${i}`,
            url: harborRegistryURLs[i],
            type: 'harbor',
            insecure: true,
            credential: {
                accessKey: getEnv('HARBOR_REGISTRY_USERNAME', 'admin'),
                accessSecret: getEnv('HARBOR_REGISTRY_PASSWORD', 'Harbor12345'),
                type: "basic"
            }
        }

        registryIds.push(harbor.createRegistry(params))
    }

    const policyIds = []
    for (let i = 0; i < policiesCount; i++) {
        const registryId = registryIds[i % registryIds.length]

        const params = {
            deletion: false,
            destRegistry: { id: registryId },
            enabled: true,
            filters: [
                { type: "name", value: `${projectName}/**` },
                { type: "tag", value: "**" },
            ],
            name: `policy-${now}-${i}`,
            override: true,
            trigger: { type: "manual" },
        }

        policyIds.push(harbor.createReplicationPolicy(params))
    }


    return {
        now,
        registryIds,
        policyIds,
    }
}

export default function ({ policyIds }) {
    const i = counter.up() - 1
    const policyId = policyIds[i % policyIds.length]

    try {
        harbor.startReplication(policyId)
    } catch (e) {
        console.log(e)
        errorRate.add(true)
    }
}

export function teardown() {
}
