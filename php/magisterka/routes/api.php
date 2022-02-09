<?php

use App\Models\Document;
use App\Models\Task;
use App\SortService;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\File;
use Illuminate\Support\Facades\Route;


Route::post('/measurement', function (Request $request) {
    return response()->json(SortService::sort($request->get('measurement'))[0]);
});

Route::post('/task', function (Request $request) {
    $task = new Task();
    $task->name = $request->get('name');
    $task->content = $request->get('content');
    $task->save();
    return response()->json(true);
});

Route::get('calc', function (Request $request) {
    return response()->json(SortService::sort(SortService::getRandomElements(5000)));
});

Route::get('/database', function (Request $request) {
    return response()->json(Document::all());
});

Route::get('/files', function (Request $request) {
    $fullText = '';

    for ($i = 1; $i < 11; $i++)
    {
        $filename = public_path() . "/Data/test$i.txt";
        $fullText .= File::get($filename);
    }

    return response()->json($fullText);
});
