// List artifact tags in the artifact which has 10K tags

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

const teardownResources = getEnv('TEARDOWN_RESOURCES', 'true') === 'true'

const artifactSize = getEnv('ARTIFACT_SIZE', '1 KiB')
const tagsCount = getEnvInt('ARTIFACT_TAGS_COUNT', '10000')

export let errorRate = new Rate('errors');

export let options = {
    setupTimeout: '24h',
    teardownTimeout: '6h',
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
    const projectName = `project-${Date.now()}`
    harbor.createProject({ projectName })

    console.log(`project ${projectName} created`)

    const repositoryName = `repository-${Date.now()}`

    const reference = harbor.prepareArtifactTags({
        projectName,
        repositoryName,
        tagsCount,
        artifactSize,
    })

    return {
        projectName,
        repositoryName,
        tagsCount,
        reference
    }
}

export default function ({ projectName, repositoryName, reference }) {
    const pageSize = 15
    const pages = Math.ceil(tagsCount / pageSize)
    const page = Math.floor(Math.random() * pages) + 1

    const params = {
        page,
        pageSize,
        withSignature: true,
        withImmutableStatus: true,
    }

    try {
        harbor.listArtifactTags(projectName, repositoryName, reference, params)
    } catch (e) {
        console.log(e)
        errorRate.add(true)
    }
}

export function teardown({ projectName }) {
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
