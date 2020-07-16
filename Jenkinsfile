pipeline {
    agent none
    environment {
        GOCACHE = '/tmp/gocache'
    }
    stages {
        stage('build') {
            agent { docker { image 'golang:1.14' } }
            steps {
                sh 'go build'
            }
        }
        stage('test') {
            agent { docker { image 'golang:1.14' } }
            steps {
                sh 'go test ./...'
            }
        }
    }
}
