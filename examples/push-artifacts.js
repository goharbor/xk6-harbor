import counter from "k6/x/counter"
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

function getEnvInt(env, def = missing) {
    return parseInt(getEnv(env, def), 10)
}

const teardownResources = getEnv('TEARDOWN_RESOURCES', 'true') === 'true'

const artifactSize = getEnv('ARTIFACT_SIZE', '10 MiB')

const contentStore = new ContentStore('artifact-push')

export let errorRate = new Rate('errors');

export let options = {
    setupTimeout: '2h',
    teardownTimeout: '1h',
    duration: '24h',
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
    const projectsCount = getEnvInt('PROJECTS_COUNT', `${options.vus}`)

    const projectNames = []
    for (let i = 0; i < projectsCount; i++) {
        const projectName = `project-${now}-${i}`
        try {
            harbor.createProject({ projectName })
            projectNames.push(projectName)
        } catch (e) {
            console.log(e)
            errorRate.add(true)
        }
    }

    const contents = contentStore.generateMany(artifactSize, options.iterations)

    return {
        projectNames,
        contents
    }
}

export default function ({ projectNames, contents }) {
    const i = counter.up() - 1

    const projectName = projectNames[i % projectNames.length]
    const blob = contents[i % contents.length]

    try {
        const repositoryName = getEnv('REPOSITORY_NAME', `repository-${Date.now()}-${i}`)

        harbor.push({
            ref: `${projectName}/${repositoryName}:tag-${i}`,
            store: contentStore,
            blobs: [blob],
        })
    } catch (e) {
        console.log(e)
        errorRate.add(true)
    }
}

export function teardown({ projectNames }) {
    contentStore.free()

    if (teardownResources) {
        for (const projectName of projectNames) {
            try {
                harbor.deleteProject(projectName, true)
            } catch (e) {
                console.log(`failed to delete project ${projectName}, error: ${e}`)
            }
        }
    } else {
        for (const projectName of projectNames) {
            console.log(`project ${projectName} keeped`)
        }
    }
}
