namespace API;

public class SortService
{
    public static IReadOnlyCollection<double> Sort(IReadOnlyCollection<double> elements)
    {
        double[] number = elements.ToArray();  
        var flag = true;
        var numLength = number.Length;  
  
        for (int i = 1; i <= numLength - 1 && flag; i++)  
        {  
            flag = false;  
            for (int j = 0; j < (numLength - 1); j++)  
            {  
                if (number[j + 1] > number[j])  
                {  
                    (number[j], number[j + 1]) = (number[j + 1], number[j]);
                    flag = true;  
                }  
            }  
        }

        return number.ToList();
    }

    public static IReadOnlyCollection<double> GetRandomElements(int size)
    {
        var random = new Random();
        var array = new List<double>();
        for (int i = 0; i < size; i++)
        {
            array.Add(random.Next(0, 10000));
        }
        return array;
    }
}