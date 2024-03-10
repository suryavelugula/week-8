pipeline {
    agent any
    
    environment {
        PATH = "$PATH:/usr/local/go/bin"
    }

    stages {
        stage('Checkout') {
            steps {
                git url: 'https://github.com/DevOps-Course-Batch-B1-24/Week8-Session-2.git', branch: 'master'
            }
        }
        
        
        stage('Build') {
            steps {
               sh 'go build'
            }
        }
        
        stage('Test') {
            steps {
                sh 'go test'
            }
        }
    }
}

