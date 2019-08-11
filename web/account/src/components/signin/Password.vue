<template>
  <v-card width="450" height="500" flat outlined>
    <v-flex mt-8 mb-5 xs12 lg12>
      <v-layout align-center justify-center>
        <v-img :src="require('../../assets/img/logo.png')" max-width="70" inline></v-img>
      </v-layout>
    </v-flex>
    <v-flex my-4 xs12 lg12>
      <h2>hatlonely hatlonely</h2>
    </v-flex>
    <v-flex my-4 xs12 lg12>
      <v-layout align-center justify-center>
        <v-card class="h-round-card px-2 py-1" outlined flat>
          <v-avatar size="20" class="mr-2" tile>
            <img :src="require('../../assets/img/hatlonely.png')" alt="avatar" />
          </v-avatar>
          <span class="body-1">hatlonely@gmail.com</span>
        </v-card>
      </v-layout>
    </v-flex>
    <v-flex my-10 mx-12>
      <v-form ref="form" v-model="valid" lazy-validation>
        <v-text-field
          label="输入您的密码"
          :append-icon="show ? 'visibility' : 'visibility_off'"
          :type="show ? 'text' : 'password'"
          @click:append="show = !show"
          :rules="[rules.required, rules.atleast8characters]"
          outlined
          filled
        ></v-text-field>
      </v-form>
    </v-flex>

    <v-flex my-12 mx-12>
      <v-layout align-left justify-center row fill-height text-left>
        <v-flex xs3>
          <v-btn text color="primary" pl-0>
            <strong>忘记密码？</strong>
          </v-btn>
        </v-flex>
        <v-flex xs6></v-flex>
        <v-flex xs3>
          <v-btn color="primary" depressed @click="validate" :disabled="!valid">下一步</v-btn>
        </v-flex>
      </v-layout>
    </v-flex>
  </v-card>
</template>

<style>
.h-round-card.v-card {
  border-radius: 9999px;
}
.h-round-card.v-card img {
  border-radius: 50%;
}
</style>

<script>
import rules from "../../assets/js/rules";

export default {
  methods: {
    validate() {
      if (this.$refs.form.validate()) {
        this.$router.push("/");
      }
    }
  },
  data() {
    return {
      show: false,
      valid: true,
      rules
    };
  }
};
</script>
