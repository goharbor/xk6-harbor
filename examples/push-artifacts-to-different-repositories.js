// Push 100k artifacts to 500 projects, 200 repositories(artifacts) per project

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
const artifactsPerProject = getEnvInt('ARTIFACTS_PER_PROJECT', '200')
const projectsCount = getEnvInt('PROJECTS_COUNT', '500')

const contentStore = new ContentStore('push-to-different-repositories')

let contents = new SharedArray('contents', function () {
    return contentStore.generateMany(artifactSize, projectsCount * artifactsPerProject)
});

export let errorRate = new Rate('errors');

export let options = {
    teardownTimeout: '1h',
    duration: '24h',
    noUsageReport: true,
    vus: artifactsPerProject,
    iterations: projectsCount * artifactsPerProject,
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
    const projectNames = []

    for (let i = 0; i < projectsCount; i++) {
        const projectName = `project-${Date.now()}-${i}`
        harbor.createProject({ projectName })
        projectNames.push(projectName)
    }

    return {
        projectNames,
    }
}

export default function ({ projectNames }) {
    const i = counter.up() - 1

    const projectName = projectNames[i % projectNames.length]
    const blob = contents[i % contents.length]

    try {
        harbor.push({
            ref: `${projectName}/repository-${Date.now()}-${i}:latest`,
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
    }
}
