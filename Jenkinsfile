pipeline {
  agent {
    docker {
      image 'golang'
    }
    
  }
  stages {
    stage('Prebuild') {
      steps {
        sh 'echo "Prebuild"'
      }
    }
    stage('Build') {
      steps {
        sh 'echo "Build"'
      }
    }
    stage('Deploy') {
      steps {
        sh 'echo "Deploy"'
      }
    }
  }
}