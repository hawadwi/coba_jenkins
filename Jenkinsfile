pipeline {
    agent any

    stages {
        stage('Test') {
            steps {
                dir('gudang-service') {
                    bat 'go test -v -cover ./...'
                }
            }
        }
    }
}
