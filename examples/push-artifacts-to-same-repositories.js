// Push 100k artifacts to one repository

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

const artifactSize = getEnv('ARTIFACT_SIZE', '1 KiB')
const artifactsCount = getEnvInt('ARTIFACTS_COUNT', '100000')

const contentStore = new ContentStore('push-to-same-repository')

let contents = new SharedArray('contents', function () {
    return contentStore.generateMany(artifactSize, artifactsCount)
});

export let errorRate = new Rate('errors');

export let options = {
    teardownTimeout: '1h',
    duration: '24h',
    noUsageReport: true,
    vus: artifactsCount > 200 ? 200 : artifactsCount,
    iterations: artifactsCount,
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

    return {
        now,
        projectName,
    }
}

export default function ({ now, projectName }) {
    const i = counter.up() - 1
    const blob = contents[i % contents.length]

    try {
        harbor.push({
            ref: `${projectName}/repository-${now}:tag-${i}`,
            store: contentStore,
            blobs: [blob],
        })
    } catch (e) {
        console.log(e)
        errorRate.add(true)
    }
}

export function teardown({ projectName }) {
    contentStore.free()

    if (teardownResources) {
        try {
            harbor.deleteProject(projectName, true)
        } catch (e) {
            console.log(`failed to delete project ${projectName}, error: ${e}`)
        }
    } else {
        console.log(`project ${projectName} keeped`)
    }
}
