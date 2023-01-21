![Frame 4](https://user-images.githubusercontent.com/59373143/176819489-a4ed02c7-67e1-4c29-8dcc-05cc503fc89e.png)


# Nashayest
Nashayest is a free and [Open API](https://en.wikipedia.org/wiki/Open_API) for detecting illegal words in a sentence to ensure that the final sentence is safe for everyone.  

<br />

## How it's works
The Nashayest database contains many illegal words in English and other languages (e.g Persian). All the illegal words found in your text will be provided to you, so you may do whatever you want with them.  

<br />

> You can see the Nashayest illegal words [here](https://nashayest.onrender.com/api_v1.0/words)

<br />

## How to use it
You can use it with `curl` or any API platform like `Postman`, or simply put it in your own app if you wish to.

```sh
$ curl https://nashayest.onrender.com/api_v1.0/?text=hello 
```

<br />

If everything goes right you should see something like this.  

```json
{
  "illegalWords":[],
  "isIllegal":false
}
```

<br />

## Todo

- [X] ~~Provide a way to add new words~~
- [X] ~~Provide a way to update/delete words~~
- [X] ~~Use a better algorithm for finding words~~

<br />

## License
The Nashayest project is under [MIT](https://github.com/sttatusx/nashayest/blob/main/LICENSE) license
