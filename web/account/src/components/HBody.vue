<template>
  <v-container>
    <v-layout mt-5 text-center wrap>
      <v-flex mb-5 xs12 lg12>
        <v-avatar size="100">
          <img :src="require('../assets/hatlonely.png')" alt="avatar" />
        </v-avatar>
      </v-flex>

      <v-flex mb-5 xs12 lg12>
        <h1>Welcome, hatlonely hatlonely</h1>
        <p>Manage your info, privacy, and security to make Google work better for you</p>
      </v-flex>

      <v-flex mb-5 xs12 md6>
        <h-card
          title="Privacy & personalization"
          description="See the data in your Google Account and choose what activity is saved to personalize your Google experience"
          image="data_personalization.png"
          linktext="Manage your data & personalization"
        />
      </v-flex>

      <v-flex mb-5 xs12 md6>
        <h-card
          title="Security issues found"
          description="Protect your account now by resolving these issues"
          image="security.png"
          linktext="Secure account"
        />
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import HCard from "./HCard.vue";

export default {
  components: {
    HCard
  },
  data: () => ({
    ecosystem: [
      {
        text: "vuetify-loader",
        href: "https://github.com/vuetifyjs/vuetify-loader"
      },
      {
        text: "github",
        href: "https://github.com/vuetifyjs/vuetify"
      },
      {
        text: "awesome-vuetify",
        href: "https://github.com/vuetifyjs/awesome-vuetify"
      }
    ],
    importantLinks: [
      {
        text: "Documentation",
        href: "https://vuetifyjs.com"
      },
      {
        text: "Chat",
        href: "https://community.vuetifyjs.com"
      },
      {
        text: "Made with Vuetify",
        href: "https://madewithvuejs.com/vuetify"
      },
      {
        text: "Twitter",
        href: "https://twitter.com/vuetifyjs"
      },
      {
        text: "Articles",
        href: "https://medium.com/vuetify"
      }
    ],
    whatsNext: [
      {
        text: "Explore components",
        href: "https://vuetifyjs.com/components/api-explorer"
      },
      {
        text: "Select a layout",
        href: "https://vuetifyjs.com/layout/pre-defined"
      },
      {
        text: "Frequently Asked Questions",
        href: "https://vuetifyjs.com/getting-started/frequently-asked-questions"
      }
    ]
  })
};
</script>
