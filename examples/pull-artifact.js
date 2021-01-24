// Pull artifact which size is 10 MiB

import harbor from 'k6/x/harbor'
import { ContentStore } from 'k6/x/harbor'

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

const conentStore = new ContentStore('artifact-pull')

export let options = {
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

export function setup() {
    harbor.initialize({
        scheme: getEnv('HARBOR_SCHEME', 'https'),
        host: getEnv('HARBOR_HOST'),
        username: getEnv('HARBOR_USERNAME', 'admin'),
        password: getEnv('HARBOR_PASSWORD', 'Harbor12345'),
        insecure: true,
    })

    const projectName = `project-${Date.now()}`
    try {
        harbor.createProject({ projectName })
    } catch (e) {
        console.log(e)
        errorRate.add(true)
    }

    harbor.push({
        ref: `${projectName}/benchmark:latest`,
        store: conentStore,
        blobs: [conentStore.generate(artifactSize)],
    })

    return {
        projectName,
    }
}

export default function ({ projectName }) {
    harbor.pull(`${projectName}/benchmark:latest`)
}

export function teardown({ projectName }) {
    conentStore.free()

    if (teardownResources) {
        try {
            harbor.deleteProject(projectName, true)
        } catch (e) {
            console.log(e)
        }
    }
}
