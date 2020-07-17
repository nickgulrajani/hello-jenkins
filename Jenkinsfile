pipeline {
    agent { docker { image 'golang:1.14' } }
    environment {
        GOCACHE = '/tmp/gocache'
    }
    stages {
        stage('build') {
            steps {
                sh 'go build'
            }
        }
        stage('build') {
            steps {
                sh 'go test ./...'
            }
        }
    }
}
