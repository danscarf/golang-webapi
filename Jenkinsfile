node('build-docker') {
    def app
    stage('Checkout') {
        // Pull the code from the repo
        checkout(
          [$class: 'GitSCM',
          branches: [[name: '*/master']],
          doGenerateSubmoduleConfigurations: false,
          extensions: [],
          submoduleCfg: [],
          userRemoteConfigs:
            [
              [
                credentialsId: 'github',
                url: 'https://github.com/danscarf/golang-webapi'
              ]
            ]
          ]
        )
    }
    stage('Build Docker image') {
        app = docker.build("dscarf/golang-webapi")
    }

    stage('Push to Docker Hub') {
        docker.withRegistry('https://registry.hub.docker.com', 'dockerhubcreds') {
            app.push("${env.BUILD_NUMBER}")
            app.push("latest")
            }
    }
}
