package main

import (
  "github.com/alloy-d/goauth"
  "log"
  "net/http"
)

func NewTwitterOAuth(key, secret, callback string) *oauth.OAuth {
  o := new(oauth.OAuth)
  o.ConsumerKey = key
  o.ConsumerSecret = secret
  o.Callback = callback + "/auth/twitter/callback"
  o.RequestTokenURL = "https://api.twitter.com/oauth/request_token"
  o.OwnerAuthURL = "https://api.twitter.com/oauth/authorize"
  o.AccessTokenURL = "https://api.twitter.com/oauth/access_token"
  o.SignatureMethod = "HMAC-SHA1"
  return o
}

func (s *Server) twitterAuthHandler(c http.ResponseWriter, req *http.Request) {
  if nickname := s.GetSession(req, "user"); nickname != "" {
    http.Redirect(c, req, ("/room/" + nickname), 302)
    return
  }

  err := s.oauth.GetRequestToken()
  if err != nil {
    log.Println(err)
    http.Redirect(c, req, "/", 302)
    return
  }

  url, err := s.oauth.AuthorizationURL()
  if err != nil {
    log.Println(err)
    http.Redirect(c, req, "/", 302)
    return
  }

  session, _ := s.cookies.Get(req, "session")
  session.Values["requestToken"] = s.oauth.RequestToken
  session.Values["requestSecret"] = s.oauth.RequestSecret
  session.Save(req, c)

  http.Redirect(c, req, url, 302)
}

func (s *Server) twitterAuthCallbackHandler(c http.ResponseWriter, req *http.Request) {
  session, _ := s.cookies.Get(req, "session")
  s.oauth.RequestToken = session.Values["requestToken"].(string)
  s.oauth.RequestSecret = session.Values["requestSecret"].(string)

  req.ParseForm()
  token := req.Form.Get("oauth_verifier")

  if err := s.oauth.GetAccessToken(token); err != nil {
    log.Println(err)
    http.Redirect(c, req, "/", 302)
    return
  }

  session.Values["user"] = s.oauth.UserName()
  session.Save(req, c)

  http.Redirect(c, req, ("/room/" + s.oauth.UserName()), 302)
}
