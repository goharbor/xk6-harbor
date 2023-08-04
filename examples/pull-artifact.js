// Pull artifact which size is 10 MiB

import { Harbor, ContentStore } from 'k6/x/harbor'
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

const artifactSize = getEnv('ARTIFACT_SIZE', '10 MiB')

const store = new ContentStore('artifact-pull')

export let successRate = new Rate('success')

export let options = {
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
    const projectName = `project-${Date.now()}`
    harbor.createProject({ projectName })

    harbor.push({
        store,
        ref: `${projectName}/benchmark:latest`,
        blobs: [store.generate(artifactSize)],
    })

    return {
        projectName,
    }
}

export default function ({ projectName }) {
    try {
        harbor.pull(`${projectName}/benchmark:latest`, { discard: false })
        successRate.add(true)
    } catch (e) {
        successRate.add(false)
        console.log(e)
    }
}

export function teardown({ projectName }) {
    store.free()

    if (teardownResources) {
        try {
            harbor.deleteProject(projectName, true)
        } catch (e) {
            console.log(e)
        }
    }
}
