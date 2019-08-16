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
            <v-menu
              ref="birthdayMenu"
              v-model="birthdayMenu"
              :close-on-content-click="false"
              transition="scale-transition"
              offset-y
              full-width
              max-width="290px"
              min-width="290px"
            >
              <template v-slot:activator="{ on }">
                <v-text-field
                  v-model="dateFormatted"
                  label="生日"
                  :rules="[rules.required]"
                  v-on="on"
                  outlined
                  filled
                  readonly
                  :value="dateFormatted = formatBirthday(birthday)"
                ></v-text-field>
              </template>
              <v-date-picker
                v-model="birthday"
                no-title
                ref="picker"
                locale="zh-cn"
                :max="new Date().toISOString().substr(0, 10)"
                min="1950-01-01"
                @change="save"
                color="primary"
                :day-format="date => date.split('-')[2]"
              ></v-date-picker>
            </v-menu>
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
        if (this.birthday == this.$store.state.account.birthday) {
          this.$router.go(-1);
          return;
        }

        this.loading = true;
        try {
          const res = await this.$store.dispatch("account/update", {
            token: this.$cookies.get("token"),
            field: "birthday",
            birthday: this.birthday
          });
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
    },
    formatBirthday(birthday) {
      if (!birthday) return null;
      const [year, month, day] = birthday.split("-");
      return `${year} 年 ${month} 月 ${day} 日`;
    },
    save(date) {
      this.$refs.birthdayMenu.save(date);
    }
  },
  watch: {
    birthday(val) {
      this.dateFormatted = this.formatBirthday(val);
    },
    birthdayMenu(val) {
      val && setTimeout(() => (this.$refs.picker.activePicker = "YEAR"));
    }
  },
  data() {
    return {
      valid: true,
      loading: false,
      dateFormatted: null,
      birthdayMenu: false,
      rules,
      birthday: this.$store.state.account.birthday
    };
  }
};
</script>

