<template>
  <v-card :loading="loading" width="450" height="500" flat outlined class="mt-12">
    <v-flex mt-8 mb-5 xs12 lg12>
      <v-layout align-center justify-center>
        <v-img :src="require('../../../assets/img/logo.png')" max-width="70" inline></v-img>
      </v-layout>
    </v-flex>
    <v-flex my-4 xs12 lg12>
      <h2>修改</h2>
    </v-flex>
    <v-flex my-4 xs12 lg12>
      <div class="body-1">您的 {{this.$config.org}} 账号</div>
    </v-flex>
    <v-flex my-10 mx-12>
      <v-form ref="form" v-model="valid" lazy-validation>
        <v-layout mx-0 row wrap>
          <v-flex xs12>
            <v-text-field
              v-model="email"
              label="邮箱"
              :rules="[rules.required, rules.validemail, rules.atmost64characters]"
              :error-messages="errors"
              outlined
              filled
              validate-on-blur
            ></v-text-field>
          </v-flex>
        </v-layout>
      </v-form>
    </v-flex>
    <v-flex my-12 mx-12>
      <v-layout align-left justify-center row fill-height text-left>
        <v-flex xs3>
          <v-btn text color="primary" @click="$router.go(-1)">返回</v-btn>
        </v-flex>
        <v-flex xs6></v-flex>
        <v-flex xs3>
          <v-btn color="primary" depressed @click="validate">保存</v-btn>
        </v-flex>
      </v-layout>
    </v-flex>
  </v-card>
</template>

<script>
const axios = require("axios");
import rules from "../../../assets/js/rules";

export default {
  methods: {
    async validate() {
      if (this.$refs.form.validate()) {
        if (this.email == this.$store.state.account.email) {
          this.$router.go(-1);
          return;
        }

        this.loading = true;
        try {
          const okTips = await this.$store.dispatch("account/verify", {
            field: "email",
            email: this.email
          });
          if (!okTips.ok) {
            this.errors = [okTips.tip];
            return;
          }

          const res = await this.$store.dispatch("account/update", {
            token: this.$cookies.get("token"),
            field: "email",
            email: this.email
          });
          console.log(res);
          if (!res.ok) {
            this.errors = [res.err];
            return;
          }
          this.$router.go(-1);
        } catch (error) {
          console.log(error);
          this.$router.push("/signin/sorry");
        } finally {
          this.loading = false;
        }
      }
    }
  },
  data() {
    return {
      valid: true,
      loading: false,
      errors: [],
      rules,
      email: this.$store.state.account.email
    };
  }
};
</script>
