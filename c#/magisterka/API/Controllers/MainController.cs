using System.Text;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;

namespace API.Controllers;

[ApiController]
[Route("api/")]
public class MainController : ControllerBase
{
    [HttpGet("files")]
    public async Task<string> GetFiles()
    {
        StringBuilder fullText = new StringBuilder();
        
        for (int i = 1; i < 11; i++)
        {
            string text = await System.IO.File.ReadAllTextAsync(Path.Combine(Environment.CurrentDirectory, "Data", $"test{i}.txt"));
            fullText.Append(text);
        }

        return fullText.ToString();
    }
    
    [HttpGet("database")]
    public async Task<IReadOnlyCollection<Document>> GetDatabase()
    {
        await using var db = new AppDbContext();
        return await db.Document.ToListAsync();
    }
    
    [HttpGet("calc")]
    public IReadOnlyCollection<double> GetCalc()
    {
        return SortService.Sort(SortService.GetRandomElements(5000));
    }
    
    [HttpPost("task")]
    public async void AddTask([FromBody] Task task)
    {
        await using var db = new AppDbContext();
        await db.Task.AddAsync(new Task{Name = task.Name, Content = task.Content});
        await db.SaveChangesAsync();
    }
    
    [HttpPost("measurement")]
    public double AddMeasurement([FromBody] IReadOnlyCollection<double> measurement)
    {
        return SortService.Sort(measurement).First();
    }
}