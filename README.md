## Daily

Daily is a crude but efficient daily function executer.

## Usage Example

  func doSomething() {
    // your function
  }
  func doSomethingElse() {
    // your function
  }
  
  secondsPastMidnight := 30
  runNow := true
  daily.Run(doSomething, secondsPastMidnight, runNow)
  daily.Run(doSomethingElse, 0, false)

The function `doSomething()` will execute 30 seconds past midnight UTC each day and it will executed right now as well.
The function `doSomethingElse()` will be executed at exactly midnight each day and will not be executed right now.
