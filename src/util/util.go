//util 包含用于处理字符串的工具函数。
package util

//Reverse 将其实参字符串以符文为单位左右反转。
func Reverse(s string) string {
	r := []rune(s)
	for i, j :=0, len(r)-1; i < len(r)/2; i, j=i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

//Sort 简单的将字符串按字符顺序排列
func Sort(s string) string {
	r := []rune(s)
	for i :=0; i < len(r)-1; i=i+1 {
		for j :=i+1; j < len(r); j=j+1 {
			if r[i] > r[j] {
				r[i], r[j] = r[j], r[i]
			}
		}
	}
	return string(r)
}

func BubbleSort(s string) string{

	values := []rune(s)
    flag:=true
    for i:=0;i<len(values)-1;i++{
        flag=true
        for j:=0;j<len(values)-i-1;j++{
            if values[j]>values[j+1]{
                values[j],values[j+1]=values[j+1],values[j]
                flag=false
            }
        }
        if flag==true{
            // 如果已经顺序对了，就不用继续冒泡排序了。
            break
        }
    }
	return string(values)
}

//QuickSort 简单的将数字按字符顺序排列
func QuickSort(values []int, left,right int) []int{
	
	temp:=values[left]
	p:=left
	i,j:=left,right 

	for i<=j{
		for j>=p && values[j]>=temp{
			j--
		}
		if j>=p{
			values[p],values[j]=values[j],values[p]
			p=j
		} 
		for i<=p && values[i]<=temp{
			i++
		}
		if i<=p{
			values[p],values[i]=values[i],values[p]
			p=i
		}
	}
	values[p]=temp
	if p-left>1{
		QuickSort(values,left,p-1)
	}
	if right-p>1{
		QuickSort(values,p+1,right)
	} 
	return values
}