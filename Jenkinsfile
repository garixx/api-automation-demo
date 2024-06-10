pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo 'Building...'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing...'
                go test
            }
        }
        stage('Deploy') {
            steps {
                echo 'No deploy...'
            }
        }
    }
}