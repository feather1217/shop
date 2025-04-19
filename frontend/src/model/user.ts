// model/user.ts
class User {
    userId: string
    displayName: string
    pictureUrl: string
  
    constructor(userId: string, displayName: string, pictureUrl: string) {
      this.userId = userId
      this.displayName = displayName
      this.pictureUrl = pictureUrl
    }
  }
  
  export default User
  