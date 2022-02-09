using Microsoft.EntityFrameworkCore;

namespace API;

public class AppDbContext : DbContext
{
    public DbSet<Task> Task { get; set; }
    public DbSet<Document> Document { get; set; }
    
    protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
    {
        var connectionString = "server=host.docker.internal;user=root;password=root;database=magisterka";
        var serverVersion = new MySqlServerVersion(new Version(8, 0, 28));
        
        optionsBuilder.UseMySql(connectionString, serverVersion);
    }
}

public class Task
{
    public int Id { get; set; }
    public string Name { get; set; }
    public string Content { get; set; }
}

public class Document
{
    public int Id { get; set; }
    public double Weight { get; set; }
    public string Name { get; set; }
    public string Content { get; set; }
}