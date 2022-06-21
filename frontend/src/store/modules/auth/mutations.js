export default {
  setToken(state, newToken) {
    state.token = newToken;
  },
  initialiseStore(state) {
    // Check if the token exists
    if(localStorage.getItem('token')) {
      this.replaceState(
          Object.assign(state, JSON.parse(localStorage.getItem('token')))
      );
    }
  }

};