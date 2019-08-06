<template>
  <v-layout mt-5 text-center wrap>
    <v-flex mb-5 xs12 lg12>
      <div class="headline">Personal info</div>
    </v-flex>
    <v-flex mb-5 xs12 lg12>
      <div class="body-1">Basic info, like your name and photo, that you use on Google services</div>
    </v-flex>

    <v-flex xs0 lg1></v-flex>
    <v-flex mb-5 xs12 lg10>
      <v-card flat outlined>
        <v-layout>
          <v-flex xs9>
            <v-card-title>Profile</v-card-title>
            <v-card-text>
              <p class="text-lg-left">
                Some info may be visible to other people using Google services.
                <a
                  href="#"
                >Learn more</a>
              </p>
            </v-card-text>
          </v-flex>
          <v-flex xs3 mt-6>
            <v-avatar size="50">
              <img :src="require('../assets/hatlonely.png')" alt="avatar" />
            </v-avatar>
          </v-flex>
        </v-layout>

        <v-list class="text-lg-left">
          <v-list-item-group>
            <v-divider class="ml-4"></v-divider>
            <v-list-item>
              <v-flex xs3>
                <div class="overline">Name</div>
              </v-flex>
              <v-list-item-title>hatlonely hatlonely</v-list-item-title>
              <v-icon small>arrow_forward_ios</v-icon>
            </v-list-item>
            <v-divider class="ml-4"></v-divider>
            <v-list-item>
              <v-flex xs3>
                <div class="overline">Birthday</div>
              </v-flex>
              <v-list-item-title>May 15, 1992</v-list-item-title>
              <v-icon small>arrow_forward_ios</v-icon>
            </v-list-item>
            <v-divider class="ml-4"></v-divider>
            <v-list-item>
              <v-flex xs3>
                <div class="overline">Gender</div>
              </v-flex>
              <v-list-item-title>Male</v-list-item-title>
              <v-icon small>arrow_forward_ios</v-icon>
            </v-list-item>
            <v-divider class="ml-4"></v-divider>
            <v-list-item>
              <v-flex xs3>
                <div class="overline">Password</div>
              </v-flex>
              <v-list-item-title>********</v-list-item-title>
              <v-icon small>arrow_forward_ios</v-icon>
            </v-list-item>
          </v-list-item-group>
        </v-list>
      </v-card>
    </v-flex>
  </v-layout>
</template>

<script>
export default {
  data: () => ({})
};
</script>
