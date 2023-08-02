// List artifacts in the repository, will prepare 100K artifacts when PROJECT_NAME or REPOSITORY_NAME env is not set by default

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
const artifactsCount = getEnvInt('ARTIFACTS_COUNT', '100000')

export let errorRate = new Rate('errors');

export let options = {
    setupTimeout: '24h',
    teardownTimeout: '6h',
    duration: '1h',
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
    let projectName = getEnv('PROJECT_NAME', '')
    let repositoryName = getEnv('REPOSITORY_NAME', '')
    let preparedInSetup = false

    if (!projectName || !repositoryName) {
        projectName = `project-${Date.now()}`

        harbor.createProject({ projectName })

        console.log(`project ${projectName} created`)

        repositoryName = `repository-${Date.now()}`

        harbor.prepareArtifacts({
            projectName,
            repositoryName,
            artifactsCount,
            artifactSize,
        })

        preparedInSetup = true
    }


    return {
        preparedInSetup,
        projectName,
        repositoryName,
    }
}

export default function ({ projectName, repositoryName }) {
    const pageSize = 15
    const pages = Math.ceil(artifactsCount / pageSize)
    const page = Math.floor(Math.random() * pages) + 1

    const params = {
        page,
        pageSize,
        withImmutableStatus: true,
        withLabel: true,
        withScanOverview: true,
        withSignature: true,
    }

    try {
        harbor.listArtifacts(projectName, repositoryName, params)
    } catch (e) {
        console.log(e)
        errorRate.add(true)
    }
}

export function teardown({ projectName, preparedInSetup }) {
    if (teardownResources && preparedInSetup) {
        try {
            harbor.deleteProject(projectName, true)
            console.log(`project ${projectName} deleted`)
        } catch (e) {
            console.log(`failed to delete project ${projectName}, error: ${e}`)
        }
    } else {
        console.log(`project ${projectName} keeped`)
    }
}
