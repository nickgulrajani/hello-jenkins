pipeline {
    agent none
    environment {
        GOCACHE = '/tmp/gocache'
    }
    stages {
        stage('build-test') {
            agent { docker { image 'golang:1.14' } }
            steps {
                sh 'go build'
                sh 'go test ./...'
            }
        }
    }
}
