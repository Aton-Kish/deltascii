# How-to guide

Prerequisites:

- [asciinema](https://asciinema.org/)

## Editing asciicast

1. recording

   ```shell
   asciinema rec ascii.cast
   ```

2. converting from asciicast to Δ-asciicast

   ```shell
   deltascii Δ -i ascii.cast -o deltascii.cast
   ```

3. editing Δ-asciicast\
   for example: correcting typos, adjusting typing speed

   <details>
   <summary>correcting typos</summary>

   ```diff
     {"version":2,"width":80,"height":24,"timestamp":1504467315,"env":{"SHELL":"/bin/zsh","TERM":"xterm-256color"}}
     [0.224325,"o","h"]
   - [0.143663,"o","w"]
   - [0.182408,"o","\b\u001b[K"]
     [0.174625,"o","e"]
     [0.160051,"o","l"]
     [0.168057,"o","l"]
     [0.215984,"o","o"]
     [0.224856,"o"," "]
     [0.208737,"o","w"]
     [0.238657,"o","o"]
     [0.187931,"o","r"]
     [0.159915,"o","l"]
     [0.192095,"o","d"]
   ```

   </details>

   <details>
   <summary>adjusting typing speed</summary>

   ```diff
     {"version":2,"width":80,"height":24,"timestamp":1504467315,"env":{"SHELL":"/bin/zsh","TERM":"xterm-256color"}}
   - [0.224325,"o","h"]
   - [0.174625,"o","e"]
   - [0.160051,"o","l"]
   - [0.168057,"o","l"]
   - [0.215984,"o","o"]
   - [0.224856,"o"," "]
   - [0.208737,"o","w"]
   - [0.238657,"o","o"]
   - [0.187931,"o","r"]
   - [0.159915,"o","l"]
   - [0.192095,"o","d"]
   + [0.2,"o","h"]
   + [0.2,"o","e"]
   + [0.2,"o","l"]
   + [0.2,"o","l"]
   + [0.2,"o","o"]
   + [0.2,"o"," "]
   + [0.2,"o","w"]
   + [0.2,"o","o"]
   + [0.2,"o","r"]
   + [0.2,"o","l"]
   + [0.2,"o","d"]
   ```

   </details>

4. re-constructing asciicast from edited Δ-asciicast

   ```shell
   deltascii Σ -i deltascii.cast -o ascii.cast
   ```

## See also

- [Command reference](./reference/README.md)
