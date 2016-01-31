# DontBeLast (in Golang)
DontBeLast is turn based 2 player game.(Human vs AI)<br>

Initial board:<br>
1<br>
11<br>
111<br>
1111<br>
11111<br>

Rules:<br>
1. The player can cut only 1's.<br>
2. After a cut, the number changes from 1 to 0 (which cannot be cut again).<br>
3. In each turn, the player can cut 1's from a single row.<br>
4. Only continuous 1's can be cut in a single turn.<br>
5. The player who cuts the last loses the game. (In the end all number changes to 0).<br>
