export interface ResponseBase {
  code: number;
  msg?: string;
}

export interface ResponseWithData<T> extends ResponseBase {
  data: T;
}

export interface UserResponse {
  id: number;
  user_name: string;
  is_admin: boolean;
  jwt_token?: string;
}

enum APIName {
  AUTH_LOGIN = "/auth/login",
}

function FetchBackend(
  input: RequestInfo | URL,
  init?: RequestInit
): Promise<any> {
  return fetch(input, {
    method: "POST",
    ...init,
    headers: {
      Authorization: localStorage.getItem("jwt") ?? "",
    },
  })
    .then((res) => res.json())
    .then((res) => {
      if (res.code === 7000 && res.msg == "Unauthorized") {
        window.location.href = "/#/login";
        return;
      }
      return res;
    });
}

export default {
  authLogin(
    username: string,
    pwd: string
  ): Promise<ResponseWithData<UserResponse>> {
    return FetchBackend(APIName.AUTH_LOGIN, {
      body: JSON.stringify({
        user_name: username,
        password: pwd,
      }),
    });
  },
};
