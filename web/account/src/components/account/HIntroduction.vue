<template>
  <v-layout mt-5 align-center justify-center row fill-height text-center wrap>
    <v-flex xs8 mt-12>
      <v-layout align-center justify-center>
        <v-img :src="require('../../assets/img/introduction.png')" max-width="400" inline></v-img>
      </v-layout>
      <v-flex my-3 xs12>
        <p>登录自己的 {{this.$config.org}} 帐号后，您可以查看和管理自己的信息、活动、安全选项和隐私偏好设置，从而让 {{this.$config.org}} 更好地为您服务。</p>
        <p>现在您可以查看和调整一些隐私选项，登录或创建帐号后还可以找到更多控制项。</p>
      </v-flex>
      <v-flex my-3 xs12>
        <v-btn color="primary" depressed to="/signin">登录</v-btn>
      </v-flex>
      <v-flex my-3 xs12>
        <v-btn text color="primary" to="/signup">创建 {{this.$config.org}} 账号</v-btn>
      </v-flex>
    </v-flex>
  </v-layout>
</template>
