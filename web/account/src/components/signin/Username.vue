<template>
  <v-card width="450" height="500" flat outlined>
    <v-flex mt-8 mb-5 xs12 lg12>
      <v-layout align-center justify-center>
        <v-img :src="require('../../assets/img/logo.png')" max-width="70" inline></v-img>
      </v-layout>
    </v-flex>
    <v-flex my-4 xs12 lg12>
      <h2>登陆</h2>
    </v-flex>
    <v-flex my-4 xs12 lg12>
      <div class="body-1">使用您的 {{this.$config.org}} 账号</div>
    </v-flex>
    <v-flex my-10 mx-12>
      <v-form ref="form" v-model="valid" lazy-validation>
        <v-text-field
          v-model="username"
          label="电子邮件或者电话号码"
          :rules="[rules.required, rules.validPhoneOrEmail]"
          :error-messages="errors"
          outlined
          filled
          validate-on-blur
        ></v-text-field>
      </v-form>
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
          <v-btn color="primary" depressed @click="validate">下一步</v-btn>
        </v-flex>
      </v-layout>
    </v-flex>
  </v-card>
</template>

<script>
const axios = require("axios");
import rules from "../../assets/js/rules";

export default {
  methods: {
    async validate() {
      if (this.$refs.form.validate()) {
        this.loading = true;
        try {
          const okTips1 = await this.$store.dispatch(
            "signin/verify",
            "username"
          );
          if (!okTips1.ok) {
            this.errors = [okTips1.tip];
            return;
          }
          this.$router.push("/signin/password");
        } catch (error) {
          this.$router.push("/signin/sorry");
        } finally {
          this.loading = false;
        }
      }
    }
  },
  computed: {
    username: {
      get() {
        return this.$store.state.signin.username;
      },
      set(username) {
        this.$store.commit("signin/setUsername", username);
      }
    }
  },
  data() {
    return {
      valid: true,
      errors: [],
      rules
    };
  }
};
</script>

