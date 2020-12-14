pipeline {
    agent any
    triggers {
        pollSCM('*/5 * * * *')
    }
    stages {
        stage("Building Images") {
            steps {
                sh "docker-compose up --detach --build"
                sh "docker image tag rblog:v0.0.23 ronaldoafonso/rblog:v0.0.23"
                sh "docker image tag rblog_nginx:v0.0.1 ronaldoafonso/rblog_nginx:v0.0.1"
                sh "docker image push ronaldoafonso/rblog:v0.0.23"
                sh "docker image push ronaldoafonso/rblog_nginx:v0.0.1"
            }
        }
    }
    post {
        always {
            sh "docker-compose down"
            mail body: "Build: ${env.BUILD_URL}", subject: 'rblog Build Status', to: 'ronaldo@vpn.ronaldoafonso.com.br'
        }
    }
}
