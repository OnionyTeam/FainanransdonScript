# FainanransdonScript
FainanransdonScript, "File Binarius Generans de 27Onion" est. Is est file binarius generans script. Hoc app files binarii generat per scripta humana-readable.

## Aedificare FainanransdonScript interpres
Directorium de project ad cie, Typus hoc signum in order versus.
```shell
go build
```
> Animadverto:
> Debes install per Golang primum.

## Syntax Overview
```
1A 1B 1C { Commentarium est. }
*3(00 0B 0C 0D 0E 3F 23) *5(00)
```
Codex superius file cum sequenti binario contento generabit:

```
1A 1B 1C 00 0B 0C 0D 0E 3F 23 00 0B 0C 0D 0E 3F 23 00 0B 0C 0D 0E 3F 23 00 00 00 00 00
```
  
Input duos digitos hexadecimales ut byte cum data hexadecimali contento outputo. Verbi gratia, sequens signum byte "61" (hexadecimal ASCII de 'a') in lima binario output scribet.
```
61
```
Forma utere sicut `*3(XX XX XX ...)`, ut fragmentum codicis inter duo brackets iterare. Verbi gratia, Sequens codicem 
```
61 62 63 *5(AB AC *3(00))
```
in output file
```
61 62 63 AB AC 00 00 00 AB AC 00 00 00 AB AC 00 00 00 AB AC 00 00 00 AB AC 00 00 00
```
binario contento output erit.

## Circa API

```go
func RunScript(script string, verbose bool = false) []byte
```
Hoc munus decurrit FinanransdonScript.  
`script`: Contentum script.  
`verbose`(Libitum): Determinare si verbosa notitia output debet. `false` valorem default est.

```go
func DecompileFromData(data []byte, verbose bool = false) string
```
Notitium ex scripta generate.  
`data`: Contentum notitia.  
`verbose`(Libitum): Determinare si verbosa notitia output debet. `false` valorem default est.
