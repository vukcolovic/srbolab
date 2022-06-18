import mutations from './mutations.js';
import getters from './getters.js';

export default {
  state() {
    return {
      userId: null,
      token: null,
      didAutoLogout: false
    };
  },
  mutations,
  getters
};