### Problem: 
Jeśli panic wystąpi wewnątrz gorutyny, to nie zabije całego programu — zatrzyma tylko tę gorutynę (nie złapaną obsługą recover).

Gorutyna wywali się z błędem, ale reszta programu może działać dalej.

Jeśli nie obsłużysz paniki, możesz mieć nieprzewidziane zachowanie i trudny do debugowania błąd.

### Rozwiązanie:

Każda gorutyna, która może panikować, powinna mieć defer z recoverem.

recover można użyć też globalnie, np. w serwerze HTTP, aby zapobiec crashowi całej aplikacji.

Bez recover panika w gorutynie często powoduje silent failure — gorutyna przestaje działać, ale nie wiadomo dlaczego