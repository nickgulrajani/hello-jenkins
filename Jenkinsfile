pipeline {
    agent { docker { image 'golang:1.14' } }
    environment {
        GOCACHE = '/tmp/gocache'
    }
    stages {
        stage('build-test') {
            steps {
                sh 'go build'
                sh 'go test ./...'
            }
        }
    }
}
