<template>
  <v-card :loading="loading" flat max-width="900" color="#fafafa">
    <v-layout mt-5 text-center wrap>
      <v-flex mb-5 xs12 lg12>
        <v-avatar size="100">
          <!-- <img :src="require('../../assets/img/hatlonely.png')" alt="avatar" />-->
          <!-- <svg width="100" height="100" :data-jdenticon-value="this.$store.state.account.email" /> -->
          <div v-html="identicon"></div>
        </v-avatar>
      </v-flex>

      <v-flex mb-5 xs12 lg12>
        <h2>欢迎使用, {{this.$store.state.account.lastName}} {{this.$store.state.account.firstName}}</h2>
      </v-flex>
      <v-flex mb-5 xs12 lg12>
        <p>管理自己的信息、隐私权和安全，从而让 {{this.$config.org}} 更好地为您服务</p>
      </v-flex>

      <v-flex my-3 xs12 md6 px-3>
        <h-card
          title="隐私权和个性化"
          description="查看您 Google 帐号中的数据，并选择保存哪些活动以用于个性化您的 Google 体验"
          image="data_personalization.png"
          linktext="管理您的数据和个性设置"
        />
      </v-flex>
      <v-flex my-3 xs12 md6 px-3>
        <h-card
          title="发现安全性问题"
          description="立即解决这些问题，保护您的账户"
          image="security.png"
          linktext="确保账户安全"
        />
      </v-flex>
      <v-flex my-3 xs12 md6 px-3>
        <h-card
          title="帐号存储空间"
          description="您的帐号存储空间由各项 Google 服务（例如 Gmail 和照片）共用"
          image="storage.png"
          linktext="管理存储空间"
        />
      </v-flex>
      <v-flex my-3 xs12 md6 px-3>
        <h-card
          title="进行隐私设置检查"
          description="此分步指南可帮助您选择适合自己的隐私设置"
          image="privacy.png"
          linktext="开始使用"
        />
      </v-flex>

      <v-flex mb-5 xs12 lg12>
        <div class="body-2">
          只有您本人可以查看您的设置。Google 致力于保护您的隐私和安全。
          <a href="#">了解详情</a>
          &nbsp;&nbsp;&nbsp;&nbsp;
          <v-avatar size="40">
            <img :src="require('../../assets/img/policy.png')" alt="avatar" />
          </v-avatar>
        </div>
      </v-flex>
    </v-layout>
  </v-card>
</template>

<script>
const axios = require("axios");
const jdenticon = require("jdenticon");
import HCard from "./HCard.vue";

export default {
  components: {
    HCard
  },
  computed: {
    identicon() {
      return jdenticon.toSvg(this.$store.state.account.email, 100);
    }
  },
  data: () => ({
    loading: false
  })
};
</script>
