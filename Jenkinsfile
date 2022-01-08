pipeline {
    agent any
    tools {
        go 'go-1.17'
    }
    environment {
        GO117MODULE = 'on'
        CGO_ENABLED = 0
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
        registry = "sanekkurt/test-jenkins"
        registryCredential = 'dockerhub'
        dockerImage = ''
    }
    stages {
//         stage('Pre Test') {
//             steps {
//                 echo 'Installing dependencies'
//                 sh 'go version'
// //                 sh 'go get -u golang.org/x/lint/golint'
//             }
//         }
//
//         stage('Test') {
//             steps {
//                 withEnv(["PATH+GO=${GOPATH}/bin"]){
//                     echo 'Running vetting'
//                     sh 'go vet ./...'
// //                     echo 'Running linting'
// //                     sh 'golint .'
//                     echo 'Running test'
//                     sh 'go test -v ./...'
//                 }
//             }
//         }
//
//         stage('Build') {
//             steps {
//                 echo 'Compiling and building'
//                 sh 'go build ./cmd/main.go'
//             }
//         }

        stage('Building our image') {
            steps {
                script {
                    dockerImage = docker.build registry + ":$BUILD_NUMBER"
                }
            }
        }

        stage('Deploy our image') {
            steps {
                script {
                    docker.withRegistry( '', registryCredential) {
                        dockerImage.push()
                    }
                }
            }
        }

        stage('Cleaning up') {
            steps {
                sh 'docker rmi $registry:$BUILD_NUMBER'
            }
        }
    }
}