# Nashayest
The Nashayest (improper) service can be used to find illegal words in a sentence (work in progress) 

<br />

# How to use it
You can use it with `curl` or any API platform like `Postman`, or simply put it in your own app if you wish to.

```sh
$ curl https://nashayest.herokuapp.com/api_v1.0/?text=hello 
```

If everything goes right you should see something like this.  

```json
{
  "illegalWords":[],
  "isIllegal":false
}
```

<br />

# License
The Nashayest project is under [MIT](https://github.com/sttatusx/nashayest/blob/main/LICENSE) license
