import { TOKEN } from "./GlobalConstants";

export const isLogin = () => {
    if (sessionStorage.getItem(TOKEN)) {
        return true;
    }

    return false;
}