<template>
  <v-card width="450" height="500" flat outlined class="mt-12">
    <v-flex mt-8 mb-5 xs12 lg12>
      <v-layout align-center justify-center>
        <v-img :src="require('../../../assets/img/logo.png')" max-width="70" inline></v-img>
      </v-layout>
    </v-flex>
    <v-flex my-4 xs12 lg12>
      <h2>出了点问题</h2>
    </v-flex>
    <v-flex my-4 xs12 lg12>
      <div class="body-1">抱歉，出了点问题，请重试</div>
    </v-flex>

    <v-flex my-12 mx-12>
      <v-layout align-left justify-center row fill-height text-left>
        <v-flex xs3>
          <v-btn text color="primary" pl-0 @click="$router.go(-1)">返回</v-btn>
        </v-flex>
      </v-layout>
    </v-flex>
  </v-card>
</template>

