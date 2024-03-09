pipeline {
    agent any

    stages {
        stage('Hello') {
            steps {
                echo 'Hello World'
            }
        }
        
        stage('go test') {
            steps {
                test 'main.go'
            }
        }
        stage('go build') {
            steps {
                build 'main.go'
            }
        }
    }
}
