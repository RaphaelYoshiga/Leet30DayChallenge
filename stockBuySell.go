package main

func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0        
    }

    sellPrice := 0;
    maxProfit:=0;

    bp:= prices[0];
    tradeProfit:= 0;
	for _, price := range prices {
        if sellPrice > price {
            bp = price;
            maxProfit += tradeProfit;
            
            tradeProfit = -1;
            sellPrice = -1;
        }else if price < bp {
            bp = price;
        } else {
            profit := price - bp;

            if profit > tradeProfit {
                tradeProfit = profit;
                sellPrice = price;
            }
        }   
    }

    if tradeProfit > 0{
        maxProfit += tradeProfit;
    }
    if maxProfit < 0{
        return 0;    
    }
    return maxProfit;
}

func max(a int, b int) int{
    if a > b {
        return a;
    }
    return b;
}