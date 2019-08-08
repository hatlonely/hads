<template>
  <v-card width="450" height="500" flat outlined>
    <v-flex mt-8 mb-5 xs12 lg12>
      <v-layout align-center justify-center>
        <v-img :src="require('../../assets/logo.png')" max-width="70" inline></v-img>
      </v-layout>
    </v-flex>
    <v-flex my-4 xs12 lg12>
      <h2>登陆</h2>
    </v-flex>
    <v-flex my-4 xs12 lg12>
      <div class="body-1">使用您的 hads 账号</div>
    </v-flex>
    <v-flex my-10 mx-12>
      <v-text-field label="电子邮件或者电话号码" outlined filled></v-text-field>
      <div class="body-2 text-left">
        <a>
          <strong>忘记了电子邮件地址？</strong>
        </a>
      </div>
    </v-flex>
    <v-flex my-10 mx-12>
      <div class="body-2 text-left">
        <p>不是您自己的计算机？请使用访客模式无痕登录。</p>
        <p>
          <a>
            <strong>了解详情</strong>
          </a>
        </p>
        <p></p>
      </div>
    </v-flex>
    <v-flex my-12 mx-12>
      <v-layout align-left justify-center row fill-height text-left>
        <v-flex xs3>
          <v-btn text color="primary" pl-0 to="/signup">创建账号</v-btn>
        </v-flex>
        <v-flex xs6></v-flex>
        <v-flex xs3>
          <v-btn color="primary" depressed to="/signin/password">下一步</v-btn>
        </v-flex>
      </v-layout>
    </v-flex>
  </v-card>
</template>
