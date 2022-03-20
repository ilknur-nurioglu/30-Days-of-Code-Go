*
 * Complete the 'solve' function below.
 *
 * The function accepts following parameters:
 *  1. DOUBLE meal_cost
 *  2. INTEGER tip_percent
 *  3. INTEGER tax_percent
 */

func solve(meal_cost float64, tip_percent int32, tax_percent int32) {
    // Write your code here
    var tip float64  = float64(tip_percent) / 100 * meal_cost
    var tax float64 = float64(tax_percent) / 100 * meal_cost
    fmt.Println(math.Round(tip+tax+meal_cost))
}
