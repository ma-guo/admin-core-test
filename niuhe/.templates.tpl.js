
// templates 注入 例子。 不能使用 js 高级语法。 
// 详细例子参考 https://handlebarsjs.com/zh/

handlebars.registerHelper('funName', function(args){
   return args;
});

handlebars.registerHelper('convertType', function (type) {
    const typeMap = {
        'string': 'string',
        'stringenum': 'string',
        'number': 'number',
        'float': 'number',
        'double': 'number',
        'int': 'number',
        'long': 'number',
        'enum': 'number',
        'integer': 'number',
        'boolean': 'boolean',
        'object': 'any',
        'array': 'any[]',
        'map': 'any',
    };
    return typeMap[type] || type;
});