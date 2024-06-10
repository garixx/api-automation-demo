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
                echo 'second'
                pwd()
            }
        }
        stage('Deploy') {
            steps {
                echo 'No deploy...'
            }
        }
    }
}