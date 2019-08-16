<template>
  <v-card width="748" height="511" flat outlined>
    <v-layout row wrap mx-0 text-left px-8>
      <v-flex xs8 pr-8>
        <v-flex mt-8 mb-5 xs12 lg12>
          <v-layout>
            <v-img :src="require('../../assets/img/logo.png')" max-width="70" inline></v-img>
          </v-layout>
        </v-flex>
        <v-flex my-4 xs12 lg12>
          <h2>隐私权及条款</h2>
        </v-flex>

        <v-flex my-4>
          <v-card class="overflow-y-auto" height="300" flat>
            <div class="body-2">
              <p>要创建 {{this.$config.org}} 帐号，您需要同意遵守下面的服务条款。</p>
              <p>此外，在您创建帐号后，我们会按照我们的隐私权政策来处理您的信息，其中包括以下这些要点：</p>

              <h4>我们在您使用 {{this.$config.org}} 服务时处理的数据</h4>
              <ul>
                <li>当您创建 {{this.$config.org}} 帐号时，我们会存储您提供的信息，如：姓名、电子邮件地址和电话号码。</li>
                <li>当您使用 {{this.$config.org}} 服务执行操作（例如在 Gmail 中写邮件或评论 YouTube 视频）时，我们会存储您创建的信息。</li>
                <li>例如，当您在 {{this.$config.org}} 地图上搜索某家餐馆或在 YouTube 上观看某个视频时，我们会处理与该活动相关的信息（其中包括您观看的视频、设备 ID、IP 地址、Cookie 数据和位置之类的信息）。</li>
                <li>此外，当您使用的应用或访问的网站采用了 {{this.$config.org}} 服务（例如，广告、Google Analytics（分析）和 YouTube 视频播放器）时，我们也会处理此类信息。</li>
              </ul>
              <p></p>

              <h4>处理数据的原因</h4>
              <p>我们之所以处理这些数据，是为了实现我们的政策中所述的目的，其中包括：</p>
              <ul>
                <li>帮助我们的服务提供更实用、更合乎需求的内容，例如更相关的搜索结果；</li>
                <li>提升我们的服务质量，并开发新的服务；</li>
                <li>
                  在 {{this.$config.org}} 服务以及与 {{this.$config.org}} 合作的网站和应用中投放个性化广告，具体取决于您的帐号设置；
                  防范欺诈和滥用行为，从而提高安全性；
                </li>
                <li>进行分析和衡量，以便了解 {{this.$config.org}} 服务的使用情况。此外，我们也有衡量 {{this.$config.org}} 服务使用情况的合作伙伴。详细了解这些特定的广告和衡量合作伙伴。</li>
              </ul>
              <p></p>

              <h4>合并数据</h4>
              <p>我们还会汇总各项 {{this.$config.org}} 服务以及您各台设备中的相关数据，以便实现上述目的。例如，我们会根据我们在您使用 {{this.$config.org}} 搜索和 YouTube 时收集的信息获知您的兴趣，从而向您展示相关广告，具体取决于您的帐号设置；我们还会利用从数以万亿计的搜索查询中收集的数据，构建拼写更正模型，供 {{this.$config.org}} 的各项服务使用。</p>

              <h4>一切由您掌控</h4>
              <p>上述某些数据可能与您的 {{this.$config.org}} 帐号相关联（具体取决于您的帐号设置），我们会将此类数据视为个人信息。现在，您只需点击下面的“更多选项”，便可控制 {{this.$config.org}} 收集和使用这类数据的方式。日后您随时可以访问我的帐号页面 (account.{{this.$config.org}}.com) 来调整控制项或撤消同意事项。</p>
            </div>
          </v-card>
        </v-flex>

        <v-flex mt-0 mx-0>
          <v-layout align-left justify-center row fill-height text-left>
            <v-flex xs3>
              <v-btn text color="primary" @click="$router.go(-1)" pl-0>取消</v-btn>
            </v-flex>
            <v-flex xs6></v-flex>
            <v-flex xs3>
              <v-btn color="primary" depressed @click="signup">我同意</v-btn>
            </v-flex>
          </v-layout>
        </v-flex>
      </v-flex>

      <v-flex xs4>
        <v-layout align-center justify-center row fill-height mx-0 class="text-center">
          <v-flex xs12>
            <v-flex xs12>
              <v-img :src="require('../../assets/img/signup_phone.svg')" max-width="244" inline></v-img>
            </v-flex>
            <v-flex xs12 px-4>
              <p class="body-1">您的个人信息不会对外公开，并且绝对安全</p>
            </v-flex>
          </v-flex>
        </v-layout>
      </v-flex>
    </v-layout>
  </v-card>
</template>

<script>
const axios = require("axios");

export default {
  methods: {
    async signup() {
      try {
        const res = await this.$store.dispatch("signup/signUp");
        this.$router.push("/account");
      } catch (error) {
        this.$router.push("/signup/sorry");
      }
    }
  }
};
</script>
