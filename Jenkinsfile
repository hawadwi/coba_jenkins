pipeline {
    agent any

    stages {

        stage('Test Gudang Service') {
            steps {
                dir('gudang-service') {
                    bat 'go test -v -coverprofile=coverage.out ./...'
                    bat 'go tool cover -func=coverage.out'
                }
            }
        }

        stage('Test Courier Service') {
            steps {
                dir('courier-service') {
                    bat 'go test -v -coverprofile=coverage.out ./...'
                    bat 'go tool cover -func=coverage.out'
                }
            }
        }

    }
}
